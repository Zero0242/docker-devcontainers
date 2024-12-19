package com.springprimerapi.controller;

import java.util.List;
import java.util.Map;

import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/notes")
public class NotesController {

    private Map<String, Object> note = Map.of(
            "id", 1,
            "title", "Sample Note",
            "content", "This is a sample note.");

    @GetMapping
    public List<Map<String, Object>> getAllNotes() {

        return List.of(
                note,
                note,
                note,
                note,
                note);
    }

    @GetMapping("/{id}")
    public Map<String, Object> getNoteById(@PathVariable Long id) {
        return note;
    }

    @PostMapping
    public Map<String, Object> createNote(@RequestBody Map<String, Object> note) {
        return note;
    }

    @PutMapping("/{id}")
    public Map<String, Object> updateNote(@PathVariable Long id, @RequestBody Map<String, Object> note) {
        return note;
    }

    @DeleteMapping("/{id}")
    public void deleteNote(@PathVariable Long id) {
        System.out.println("Done");
    }
}