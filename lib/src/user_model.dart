class User {
  String id;
  String email;
  String password;

  User(this.id, this.email, this.password);

  factory User.fromJson(Map<String, dynamic> user) =>
      User(user['id'], user['email'], user['password']);

  Map toJson() => {'id': id, 'email': email, 'password': password};
}
