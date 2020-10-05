class Consent {
  final String token;
  final Map<String, bool> claims;

  Consent(this.token, this.claims);

  Map<String, dynamic> toJson() {
    return {
      "Token": this.token,
      "Claims": this.claims,
    };
  }
}
