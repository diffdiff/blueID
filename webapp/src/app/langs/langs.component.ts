import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { LangService } from '../lang.service';
import Lang from '../lang';

@Component({
  selector: 'app-langs',
  templateUrl: './langs.component.html',
  styleUrls: ['./langs.component.css']
})
export class LangsComponent implements OnInit {

  Langs: Lang[];
  constructor(private route: ActivatedRoute, private router: Router, private ls: LangService) { }

  ngOnInit() {
    this.getLangs();
  }

  getLangs() {
    this.ls.getLangs().subscribe(data => {
      this.Langs = data;
    });
  }
}
