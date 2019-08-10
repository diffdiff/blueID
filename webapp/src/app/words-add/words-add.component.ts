import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { WordService } from '../word.service';
import Lang from '../lang';
import { LangService } from '../lang.service';

@Component({
  selector: 'app-words-add',
  templateUrl: './words-add.component.html',
  styleUrls: ['./words-add.component.css']
})
export class WordsAddComponent implements OnInit {

  Langs: Lang[];
  angForm: FormGroup;

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private ws: WordService, private ls: LangService) { }

  ngOnInit() {
    this.getLangs();
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      WordText: ['', Validators.required],
      Language: [''],
      Example: [''],
      Description: ['']
    });
  }

  getLangs() {
    this.ls.getLangs().subscribe(langs => {
      this.Langs = langs;
    });
  }

  addWord(wordText, langID, example, description) {
    this.ws.addWord(wordText, langID, example, description).subscribe(res => {
      this.router.navigate(['words/index']);
    });
  }
}
