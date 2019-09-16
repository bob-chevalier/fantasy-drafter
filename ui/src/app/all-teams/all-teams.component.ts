import { Component, OnInit } from '@angular/core';
import { TEAMS } from '../mock-teams';

@Component({
  selector: 'app-all-teams',
  templateUrl: './all-teams.component.html',
  styleUrls: ['./all-teams.component.css']
})
export class AllTeamsComponent implements OnInit {
  teams = TEAMS;

  constructor() {
  }

  ngOnInit() {
  }

}
