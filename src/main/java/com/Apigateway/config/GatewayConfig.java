//package com.Apigateway.config;
//
//import org.springframework.beans.factory.annotation.Value;
//import org.springframework.cloud.gateway.route.RouteLocator;
//import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
//import org.springframework.context.annotation.Bean;
//import org.springframework.context.annotation.Configuration;
//
//@Configuration
//public class GatewayConfig {
//
//    @Value("")
//
//    @Bean
//    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
//        return builder.routes()
//
//                .route("user-service", r -> r.path("/api/v1/user/**")
//                        .uri("http://localhost:8081"))
//
//                .route("auth-service", r -> r.path("/api/v1/auth/**")
//                        .uri("http://localhost:8082"))
//
//                .route("other-service", r -> r.path("/api/v1/other/**")
//                        .uri("http://localhost:8083"))
//
//
//                .build();
//    }
//}