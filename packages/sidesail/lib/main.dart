import 'dart:async';

import 'package:flutter/material.dart';
import 'package:sidesail/app.dart';
import 'package:sidesail/config/dependencies.dart';
import 'package:sidesail/config/runtime_args.dart';
import 'package:sidesail/config/sidechains.dart';

Future<void> start() async {
  WidgetsFlutterBinding.ensureInitialized();

  String appName;
  Sidechain chain;
  // Sanity check we're getting a supported chain.
  switch (RuntimeArgs.chain.toLowerCase()) {
    case '': // default to testchain
    case 'testchain':
      chain = TestSidechain();
      appName = chain.name;
      break;

    case 'ethside':
      chain = EthereumSidechain();
      appName = chain.name;
      break;

    case 'zside':
      chain = ZCashSidechain();
      appName = chain.name;
      break;

    default:
      return runApp(
        SailApp(
          builder: (context, router) => Center(
            child: Text(
              'Unsupported chain: ${RuntimeArgs.chain}',
              style: const TextStyle(fontSize: 40),
            ),
          ),
        ),
      );
  }

  await initDependencies(chain);
  runApp(
    SailApp(
      // the initial route is defined in routing/router.dart
      builder: (context, router) => MaterialApp.router(
        routerDelegate: router.delegate(),
        routeInformationParser: router.defaultRouteParser(),
        title: appName,
        theme: ThemeData(
          fontFamily: 'SourceCodePro',
          colorScheme: ColorScheme.fromSwatch().copyWith(secondary: const Color(0xffFF8000)),
        ),
      ),
    ),
  );
}

void main() {
  // the application is launched function because some startup things
  // are async
  start();
}
