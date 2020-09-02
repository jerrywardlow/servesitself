# This is a basic web server written in Python which serves it's own source code

import http.server
import socketserver

class sourceHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.path = 'server.py'
        return http.server.SimpleHTTPRequestHandler.do_GET(self)

handler = sourceHandler

httpd = socketserver.TCPServer(("", 8080), handler)

print("Server running on port 8080...")

try:
    httpd.serve_forever()
except (KeyboardInterrupt, SystemExit):
    print("Terminating...")
