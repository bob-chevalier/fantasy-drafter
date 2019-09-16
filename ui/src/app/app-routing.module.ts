import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TeamComponent } from './team/team.component';
import { AllTeamsComponent } from './all-teams/all-teams.component';
import { PlayersComponent } from './players/players.component';
import { DraftBoardComponent } from './draft-board/draft-board.component';

const routes: Routes = [
    { path: '', redirectTo: '/draft-board', pathMatch: 'full' },
    { path: 'draft-board', component: DraftBoardComponent },
    { path: 'team', component: TeamComponent },
    { path: 'players', component: PlayersComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
