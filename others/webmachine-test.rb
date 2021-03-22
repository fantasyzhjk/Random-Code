require 'webmachine'

class MyResource < Webmachine::Resource
  def to_html
    "<html><body>Hello, world!</body></html>"
  end
end

# Start a web server to serve requests via localhost
MyResource.run