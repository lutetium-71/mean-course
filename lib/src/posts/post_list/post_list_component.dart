import 'package:angular/angular.dart';

import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import '../../post_model.dart';
import '../../post_list_service.dart';

@Component(
  selector: 'app-post-list',
  styleUrls: ['post_list_component.css'],
  templateUrl: 'post_list_component.html',
  directives: [
    MaterialExpansionPanel,
    NgFor,
    NgIf,
  ],
  providers: [ClassProvider(PostListService)],
)
class PostListComponent implements OnInit {
  final PostListService postListService;

  List<Post> postList = [];

  PostListComponent(this.postListService);
  var controller = new StreamController<String>();
  @override
  Future<Null> ngOnInit() async {
    this.postList = await this.postListService.getPostList();
    this
        .postListService
        .getPostUpdateListener()
        .stream
        .listen((List<Post> posts) => {});
  }
}
