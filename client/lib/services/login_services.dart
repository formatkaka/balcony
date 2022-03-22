import 'dart:convert';

import 'package:client/models/login_models.dart';
import 'package:http/http.dart' as http;

Future<SendOTP> sendOTP(String mobileNum) async {
  final response = await http.get(
    Uri.parse("http://192.168.0.102:3000/otp?mobile_num=$mobileNum"),
  );

  if (response.statusCode == 200) {
    return SendOTP.fromJSON(jsonDecode(response.body));
  }

  throw Exception('Failed to Send OTP');
}

Future<dynamic> verifyOtpAndLogin(String mobileNum, String otp) async {
  final response = await http.post(
    Uri.parse("http://192.168.0.102:3000/login"),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(<String, String>{'mobile_num': mobileNum, 'otp': otp}),
  );

  if (response.statusCode == 200) {
    return VerifyOTP.fromJSON(jsonDecode(response.body));
  }

  throw Exception("Failed to verify OTP");
}
