class Post {
  String id;
  String title;
  String content;

  Post(this.id, this.title, this.content);

  factory Post.fromJson(Map<String, dynamic> post) =>
      Post(post['id'], post['title'], post['content']);

  Map toJson() => {'id': id, 'title': title, 'content': content};
}
