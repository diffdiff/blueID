import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LangService } from './lang.service';
import { HttpClientModule } from '@angular/common/http';
import { LangAddComponent } from './lang-add/lang-add.component';
import { LangsComponent } from './langs/langs.component';
import { WordsComponent } from './words/words.component';
import { WordsAddComponent } from './words-add/words-add.component';
import { TranslationsComponent } from './translations/translations.component';
import { TranslationAddComponent } from './translation-add/translation-add.component';

@NgModule({
  declarations: [
    AppComponent,
    LangAddComponent,
    LangsComponent,
    WordsComponent,
    WordsAddComponent,
    TranslationsComponent,
    TranslationAddComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [LangService],
  bootstrap: [AppComponent]
})
export class AppModule { }
