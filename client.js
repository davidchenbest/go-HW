const net = require('net');

const client = new net.Socket();

client.connect(8080, 'localhost', () => {
  console.log('Connected to TCP server');
  
  // Send a message to the server
  client.write('Hello, server!\n');
});

client.on('data', (data) => {
  console.log('Received:', data.toString());
});

client.on('close', () => {
  console.log('Connection closed');
});
