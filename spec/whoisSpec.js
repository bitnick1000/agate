var Whois = require('../whois');
describe('whois', function() {
	it('should be able to lookup domain name', function(done) {
		var whois1 = new Whois;
		whois1.lookup('google.com', function(err, isRegistered) {
			expect(err).toBeNull();
			expect(isRegistered).toBeTruthy();
			done();
		});
	});
	describe('google.com', function() {
		var whois2 = new Whois;
		it('should be registered', function(done) {
			whois2.lookup('google.com', function(err, isRegistered) {
				expect(err).toBeNull();
				expect(isRegistered).toBeTruthy();
				done();
			});
		});
	});
	describe('domain-name-eb3.com', function() {
		var whois3 = new Whois;
		it('should not be registered', function(done) {
			whois3.lookup('domain-name-eb3.com', function(err, isRegistered) {
				expect(err).toBeNull();
				expect(isRegistered).toBeFalsy();
				done();
			});
		});
	});
});