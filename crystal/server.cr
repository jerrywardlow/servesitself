# This is a basic web server written in Crystal which serves it's own source code

require "http/server"

# Handle SIGINT/SIGTERM
Signal::INT.trap { puts "Caught SIGINT, bye!"; exit }
Signal::TERM.trap { puts "Caught SIGTERM, sorry!"; exit }

server = HTTP::Server.new do |context|
    context.response.content_type = "text/plain"
    context.response.print "Placeholder"
end

address = server.bind_tcp "0.0.0.0", 8080
puts "Server running on port 8080..."
server.listen
