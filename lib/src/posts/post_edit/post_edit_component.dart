import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_components/focus/focus.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/material_input/material_input.dart';

import '../../post_model.dart';
import '../../post_list_service.dart';
import '../../route_paths.dart';

@Component(
  selector: 'app-post-create',
  styleUrls: [
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
    'post_edit_component.css'
  ],
  templateUrl: 'post_edit_component.html',
  directives: [
    AutoFocusDirective,
    formDirectives,
    MaterialButtonComponent,
    MaterialIconComponent,
    materialInputDirectives,
    NgForm,
    NgIf,
  ],
)
class PostEditComponent implements OnActivate {
  final PostListService _postListService;
  final Router _router;
  String postId;
  Post post;
  Post postUpdate;

  PostEditComponent(this._postListService, this._router);

  @override
  void onActivate(_, RouterState current) async {
    postId = current.parameters['id'];
    postUpdate = await _postListService.getPost(postId);
  }

  onEditPost(NgForm form) {
    Post post = Post(postId, form.value["title"], form.value["content"]);
    _postListService.updatePost(post);
    _gotoHome();
  }

  Future<NavigationResult> _gotoHome() => _router.navigate(_homeUrl());

  String _homeUrl() => RoutePaths.posts.toUrl();
}
