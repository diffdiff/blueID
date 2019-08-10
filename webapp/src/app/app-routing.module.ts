import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LangAddComponent } from './lang-add/lang-add.component';
import { LangsComponent } from './langs/langs.component';
import { WordsComponent } from './words/words.component';
import { WordsAddComponent } from './words-add/words-add.component';
import { TranslationsComponent } from './translations/translations.component';
import { TranslationAddComponent } from './translation-add/translation-add.component';


const routes: Routes = [
  {
    path: 'lang',
    component: LangAddComponent
  },
  {
    path: 'lang/index',
    component: LangsComponent
  },
  {
    path: 'words/index',
    component: WordsComponent
  },
  {
    path: 'words/add',
    component: WordsAddComponent
  },
  {
    path: 'translations/index',
    component: TranslationsComponent
  },
  {
    path: 'translation/add',
    component: TranslationAddComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
