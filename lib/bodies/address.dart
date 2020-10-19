class Address {
  final num streetNo;
  final String street;
  final String unitNo;
  final String estateName;
  final String suburb;
  final String city;
  final String province;
  final String postalCode;
  final String coordinates;
  final bool isDelivery;

  Address(
      this.streetNo,
      this.street,
      this.unitNo,
      this.estateName,
      this.suburb,
      this.city,
      this.province,
      this.postalCode,
      this.coordinates,
      this.isDelivery);

  Map<String, dynamic> toJson() {
    return {
      "StreetNo": this.streetNo,
      "Street": this.street,
      "UnitNo": this.unitNo,
      "EstateName": this.estateName,
      "Suburb": this.suburb,
      "City": this.city,
      "Province": this.province,
      "PostalCode": this.postalCode,
      "Coordinates": this.coordinates,
      "IsDelivery": this.isDelivery
    };
  }
}
