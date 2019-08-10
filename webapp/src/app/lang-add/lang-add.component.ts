import { Component, OnInit } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { LangService } from '../lang.service';

@Component({
  selector: 'app-lang-add',
  templateUrl: './lang-add.component.html',
  styleUrls: ['./lang-add.component.css']
})
export class LangAddComponent implements OnInit {

  angForm: FormGroup;

  constructor(private route: ActivatedRoute, private router: Router, private fb: FormBuilder, private ls: LangService) {
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      LangText: ['', Validators.required ]
    });
  }

  addLang(langText: string) {
    this.route.params.subscribe(params => {
      this.ls.createLang(langText);
      this.router.navigate(['lang/index']);
    });
  }

  ngOnInit() {
  }
}
