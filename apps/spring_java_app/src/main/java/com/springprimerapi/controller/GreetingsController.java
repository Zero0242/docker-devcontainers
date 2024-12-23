package com.springprimerapi.controller;

import java.util.Map;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/demo")
public class GreetingsController {
    @GetMapping("/greeting")
    @ResponseBody
    public String greeting() {
        return "Hello, World!";
    }

    @GetMapping("/greeting/param/{name}")
    @ResponseBody
    public Map<String, Object> greetingWithName(@PathVariable String name) {
        return Map.of(
                "ok", true,
                "status", 200,
                "message", "Hello, " + name + "!");
    }

    @GetMapping("/greeting/query")
    @ResponseBody
    public Map<String, Object> greetingWithQuery(
            @RequestParam(name = "name", defaultValue = "Spring") String queryName) {
        return this.greetingWithName(queryName);
    }

    @GetMapping("/error")
    @ResponseBody
    public String error() {
        return "An error has occurred!";
    }
}
