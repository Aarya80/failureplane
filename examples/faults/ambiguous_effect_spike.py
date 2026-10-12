#!/usr/bin/env python3
"""Phase-0 loopback-only demonstration: an external write can succeed even
when its caller receives no response. No FailurePlane server is involved.
Requires only Python 3 standard library; makes no outside network requests.
"""
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from threading import Thread
from urllib.request import Request, urlopen
from http.client import RemoteDisconnected
import json
import socket

ledger = {}
first_response_dropped = set()

class Provider(BaseHTTPRequestHandler):
    def do_POST(self):
        key = self.headers.get('Idempotency-Key', '')
        if not key:
            self.send_error(400, 'idempotency key required')
            return
        length = int(self.headers.get('Content-Length', 0))
        body = self.rfile.read(length)
        if key not in ledger:
            ledger[key] = {'key': key, 'body': body.decode(), 'accepted': True}
        if key not in first_response_dropped:
            first_response_dropped.add(key)
            self.close_connection = True
            try:
                self.connection.shutdown(socket.SHUT_RDWR)
            except OSError:
                pass
            self.connection.close()
            return
        payload = json.dumps(ledger[key]).encode()
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Content-Length', str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)

    def log_message(self, format, *args):
        pass  # keep sample output deterministic and short

server = ThreadingHTTPServer(('127.0.0.1', 0), Provider)
thread = Thread(target=server.serve_forever, daemon=True)
thread.start()
try:
    url = f'http://127.0.0.1:{server.server_port}/charge'
    key = 'order-481-charge'
    request = Request(url, data=b'{"amount":12900}', method='POST', headers={'Idempotency-Key': key})
    try:
        urlopen(request, timeout=3).read()
        raise AssertionError('expected the response to be deliberately lost')
    except (RemoteDisconnected, ConnectionResetError):
        print('Client: lost response, outcome UNKNOWN')

    assert len(ledger) == 1 and ledger[key]['accepted']
    print('Provider: write was actually accepted before response loss')

    # This retry is safe ONLY because the mock provider explicitly implements
    # idempotency for the same key, not because the client can infer outcome.
    with urlopen(request, timeout=3) as response:
        assert response.status == 200
        assert json.loads(response.read())['key'] == key
    assert len(ledger) == 1
    print('Provider-idempotent retry: one recorded write, not two')
    print('LESSON: no blind retry for non-idempotent downstream effects.')
finally:
    server.shutdown()
    server.server_close()
    thread.join(timeout=3)
