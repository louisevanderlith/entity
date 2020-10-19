class Role {
  final String profileID;
  final num role;

  Role(this.profileID, this.role);

  Map<String, dynamic> toJson() {
    return {
      "ProfileID": this.profileID,
      "Role": this.role,
    };
  }
}
