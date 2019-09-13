import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RosterComponent } from './roster/roster.component';
import { PlayersComponent } from './players/players.component';

const routes: Routes = [
    { path: '', redirectTo: '/roster', pathMatch: 'full' },
    { path: 'roster', component: RosterComponent },
    { path: 'players', component: PlayersComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
