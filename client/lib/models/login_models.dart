class SendOTP {
  final String response;
  final String otp;

  const SendOTP({required this.otp, required this.response});

  factory SendOTP.fromJSON(Map<String, dynamic> json) {
    return SendOTP(otp: json['otp'], response: json['err']);
  }
}

class VerifyOTP {
  final String response;
  final String token;

  const VerifyOTP({required this.response, required this.token});

  factory VerifyOTP.fromJSON(Map<String, dynamic> json) {
    return VerifyOTP(response: json['response'], token: json['token']);
  }
}
