import 'package:flutter/material.dart';

class FormField extends StatelessWidget {
  const FormField(
      {required this.hintText,
      required this.callback,
      required this.customValidator,
      Key? key})
      : super(key: key);

  final String hintText;
  final Function(int) callback;
  final Function(String?) customValidator;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      decoration: InputDecoration(hintText: hintText),
      onChanged: (value) {
        callback(int.parse(value));
      },
      validator: (value) {
        return customValidator(value);
      },
    );
  }
}
