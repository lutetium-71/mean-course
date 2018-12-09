import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/material_input/material_input.dart';

import '../../models/post.dart';

@Component(
    selector: 'app-post-create',
    styleUrls: [
      'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
      'post_create_component.css'
    ],
    templateUrl: 'post_create_component.html',
    directives: [
      formDirectives,
      MaterialButtonComponent,
      MaterialIconComponent,
      materialInputDirectives,
      NgForm,
      NgIf,
    ])
class PostCreateComponent {
  String enteredTitle = '';
  String enteredContent = '';

  final _postCreated = new StreamController<Post>();

  void onAddPost(NgForm form) {
    Post post = Post(form.value["title"], form.value["content"]);
    _postCreated.add(post);
  }

  bool checkEntry() => enteredTitle.isEmpty || enteredContent.isEmpty;

  @Output()
  Stream<Post> get postCreated => _postCreated.stream;
}
