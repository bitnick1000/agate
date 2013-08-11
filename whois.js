var net = require('net');
function whois() {
	this.client = new net.Socket();
	this.isRegistered = true;
	this.HOST = 'whois.crsnic.net';
	this.PORT = 43;
}

whois.prototype.lookup = function(domainName, callback) {
	var self = this;

	self.client.on('connect', function() {
		// console.log('on connect ' + self.client.remoteAddress);
	});

	self.client.on('data', function(data) {
		// console.log('on data');
		if (-1 != data.toString().indexOf('No match')) {
			self.isRegistered = false;
		}
		//console.log('DATA: ' + data.toString());
	});

	self.client.on('end', function() {
		// console.log('socket on end');
	});

	self.client.on('timeout', function() {
		console.log('socket on timeout');
	});

	self.client.on('drain', function() {
		console.log('on drain');
	});

	self.client.on('error', function() {
		// The 'close' event will be called directly following this event.
		console.log('socket on error');
	});

	self.client.on('close', function(had_error) {
		// console.log('on colose with error ? ' + had_error);
		if (had_error) {
			callback(1);
			//console.log('emit done');
			//self.emit('done', null, self.isRegistered);
		} else {
			callback(null, self.isRegistered);
		}
		//module.exports.emit("successed", "test", "test2");
	});
	self.client.connect(self.PORT, self.HOST, function() {
		self.client.write(domainName + '\r\n', 'ascii', function() {
			// console.log('write successed');
		});
	});
}

module.exports = whois;