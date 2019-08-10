import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TranslationService } from '../translation.service';

@Component({
  selector: 'app-translations',
  templateUrl: './translations.component.html',
  styleUrls: ['./translations.component.css']
})
export class TranslationsComponent implements OnInit {

  Translations: any[];

  constructor(private route: ActivatedRoute, private router: Router, private ts: TranslationService) { }

  ngOnInit() {
    this.getTranslations();
  }

  getTranslations() {
    this.ts.getTranslations().subscribe(data => {
      this.Translations = data;
    });
  }

  deleteTranslation(translationID) {
    this.ts.deleteTranslation(translationID).subscribe(_ => {
      this.Translations = this.Translations.filter(t => t.ID !== translationID);
    });
  }

}
