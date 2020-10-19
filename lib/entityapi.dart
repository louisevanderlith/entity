import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/entity.dart';

Future<HttpRequest> updateEntity(Key key, Entity obj) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/${key.toJson()}";
  final data = jsonEncode(obj);

  return invokeService("PUT", url, data);
}

Future<HttpRequest> createEntity(Entity obj) async {
  var apiroute = getEndpoint("entity");
  var url = "${apiroute}/";
  final data = jsonEncode(obj);

  return invokeService("POST", url, data);
}
