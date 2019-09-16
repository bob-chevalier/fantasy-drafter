import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TeamComponent } from './team/team.component';
import { AllTeamsComponent } from './all-teams/all-teams.component';
import { PlayersComponent } from './players/players.component';

const routes: Routes = [
    { path: '', redirectTo: '/all-teams', pathMatch: 'full' },
    { path: 'all-teams', component: AllTeamsComponent },
    { path: 'team', component: TeamComponent },
    { path: 'players', component: PlayersComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
