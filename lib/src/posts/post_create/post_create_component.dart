import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/material_input/material_input.dart';

import '../../post_model.dart';
import '../../post_list_service.dart';

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
  ],
)
class PostCreateComponent {
  final PostListService postListService;

  PostCreateComponent(this.postListService);

  onAddPost(NgForm form) {
    Post post = Post(form.value["title"], form.value["content"]);
    this.postListService.addPost(post);
    form.reset();
  }
}
