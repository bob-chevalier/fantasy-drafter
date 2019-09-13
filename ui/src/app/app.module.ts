import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { RosterComponent } from './roster/roster.component';
import { AppRoutingModule } from './app-routing.module';
import { PlayersComponent } from './players/players.component';

@NgModule({
  declarations: [
    AppComponent,
    RosterComponent,
    PlayersComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
