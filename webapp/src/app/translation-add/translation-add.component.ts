import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TranslationService } from '../translation.service';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import Word from '../word';
import { WordService } from '../word.service';

@Component({
  selector: 'app-translation-add',
  templateUrl: './translation-add.component.html',
  styleUrls: ['./translation-add.component.css']
})
export class TranslationAddComponent implements OnInit {

  Words: Word[];
  angForm: FormGroup;
  constructor(private route: ActivatedRoute, private router: Router,
              private fb: FormBuilder, private ts: TranslationService, private ws: WordService) { }

  ngOnInit() {
    this.createForm();
    this.getWords();
  }

  createForm() {
    this.angForm = this.fb.group({
      Word1: [],
      Word2: []
    });
  }

  getWords() {
    this.ws.getAllWords().subscribe(words => {
      this.Words = words;
    });
  }

  addTranslation(wordID1: number, wordID2: number) {
    console.log(wordID1, wordID2)
    this.ts.addTranslation(wordID1, wordID2).subscribe(res => {
      console.log('Done');
      this.router.navigate(['translations/index']);
    });
  }

}
