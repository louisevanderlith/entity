import 'address.dart';
import 'user.dart';

class Entity {
  final String name;
  final String profileKey;
  final User user;
  final String identification;
  final List<Address> addresses;

  Entity(this.name, this.profileKey, this.user, this.identification,
      this.addresses);

  Map<String, dynamic> toJson() {
    return {
      "Name": this.name,
      "ProfileKey": this.profileKey,
      "User": this.user,
      "Identification": this.identification,
      "Addresses": this.addresses
    };
  }
}
