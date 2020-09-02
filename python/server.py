import http.server
import socketserver

class sourceHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.path = 'server.py'
        return http.server.SimpleHTTPRequestHandler.do_GET(self)

handler = sourceHandler

httpd = socketserver.TCPServer(("", 8080), handler)

print("serving at port 8080...")

httpd.serve_forever()
