import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:posts/app_component.template.dart' as ng;
import 'package:http/browser_client.dart';

import 'main.template.dart' as self;

@GenerateInjector([
  routerProviders,
  ClassProvider(BrowserClient),
])
final InjectorFactory injector = self.injector$Injector;
void main() {
  runApp(ng.AppComponentNgFactory, createInjector: injector);
}
