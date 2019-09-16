import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { TeamComponent } from './team/team.component';
import { AppRoutingModule } from './app-routing.module';
import { PlayersComponent } from './players/players.component';
import { AllTeamsComponent } from './all-teams/all-teams.component';
import { DraftBoardComponent } from './draft-board/draft-board.component';

@NgModule({
  declarations: [
    AppComponent,
    TeamComponent,
    PlayersComponent,
    AllTeamsComponent,
    DraftBoardComponent
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
