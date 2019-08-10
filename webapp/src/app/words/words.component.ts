import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { WordService } from '../word.service';
import Word from '../word';

@Component({
  selector: 'app-words',
  templateUrl: './words.component.html',
  styleUrls: ['./words.component.css']
})
export class WordsComponent implements OnInit {

  Words: Word[];
  constructor(private route: ActivatedRoute, private router: Router, private ws: WordService) { }

  ngOnInit() {
    this.getAllWords();
  }

  getAllWords() {
    this.ws.getAllWords().subscribe(words => {
      this.Words = words;
    });
  }

}
