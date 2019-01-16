import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import '../routes.dart';

@Component(
  selector: 'app-header',
  styleUrls: [
    'package:angular_components/app_layout/layout.scss.css',
    'header_component.css'
  ],
  templateUrl: 'header_component.html',
  directives: [routerDirectives],
  exports: [RoutePaths, Routes],
)
class HeaderComponent {}
