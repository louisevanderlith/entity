import 'package:mango_secure/bodies/contact.dart';

import 'role.dart';

class User {
  final String name;
  final bool verified;
  final String email;
  final String password;
  final List<Contact> contacts;
  final List<String> resources;
  final List<Role> roles;

  User(this.name, this.verified, this.email, this.password, this.contacts,
      this.resources, this.roles);

  Map<String, dynamic> toJson() {
    return {
      "Name": this.name,
      "Verified": this.verified,
      "Email": this.email,
      "Password": this.password,
      "Contacts": this.contacts,
      "Resources": this.resources,
      "Roles": this.roles
    };
  }
}
