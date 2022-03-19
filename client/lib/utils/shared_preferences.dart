import 'package:shared_preferences/shared_preferences.dart';

void saveToSharedPref(String key, String value) async {
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString(key, value);
}
