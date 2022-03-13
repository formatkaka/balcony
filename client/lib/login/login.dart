import 'package:flutter/material.dart';

import 'field.dart' as FormFieldCustom;

class LoginForm extends StatefulWidget {
  const LoginForm({Key? key}) : super(key: key);

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  String buttonState = 'OTP';
  String buttonText = 'Send OTP';
  int mobileNum = 0;
  int Otp = 0;

  void buttonPressed() {
    if (buttonState == 'OTP') {
      // sendOTP

      if (_formKey.currentState!.validate()) {
        // Process data.
      }

      setState(() {
        buttonState = 'CHECK';
        buttonText = 'Verify OTP';
      });
    } else {
      // verify OTP
      print(mobileNum);
      print(Otp);
    }
  }

  void setMobileNum(int mobileNumX) {
    setState(() {
      mobileNum = mobileNumX;
    });
  }

  void setOTP(int OtpX) {
    setState(() {
      Otp = OtpX;
    });
  }

  dynamic numValidator(int length) {
    dynamic innerFunc(String? value) {
      if (value != null && value != '' && value!.length == length) {
        print("Entered");
        var parseVal = int.tryParse(value!);
        if (parseVal != null) return null;
      }

      return 'Incorrect';
    }

    return innerFunc;
  }

  @override
  Widget build(BuildContext context) {
    return Form(
        key: _formKey,
        child: Column(
          children: [
            FormFieldCustom.FormField(
                hintText: "Enter phone number",
                callback: setMobileNum,
                customValidator: numValidator(10)),
            if (buttonState == 'CHECK')
              FormFieldCustom.FormField(
                  hintText: "Enter OTP",
                  callback: setOTP,
                  customValidator: numValidator(4)),
            ElevatedButton(onPressed: buttonPressed, child: Text(buttonText))
          ],
        ));
  }
}
