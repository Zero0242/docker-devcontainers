package com.springprimerapi.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController {

    @GetMapping("/index.html")
    public String index() {
        return "Your First Return";
    }

}
