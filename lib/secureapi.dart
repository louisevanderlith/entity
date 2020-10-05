import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/requester.dart';

import 'bodies/consent.dart';
import 'bodies/login.dart';
import 'bodies/register.dart';

Future<HttpRequest> sendLogin(Login obj) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/login";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> sendForgot(String identity) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/forgot";
  final data = jsonEncode(identity);

  return invokeService("POST", url, data);
}

Future<HttpRequest> sendRegister(Register obj) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/register";
  final data = jsonEncode(obj.toJson());

  return invokeService("POST", url, data);
}

Future<HttpRequest> sendConsent(Consent obj) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/consent";
  final data = jsonEncode(obj);

  return invokeService("POST", url, data);
}