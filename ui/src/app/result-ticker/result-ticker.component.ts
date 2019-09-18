import { Component, OnInit } from '@angular/core';
import { DraftPick } from '../draft-pick';
import { DRAFT_PICKS } from '../mock-draft-picks';

@Component({
  selector: 'app-result-ticker',
  templateUrl: './result-ticker.component.html',
  styleUrls: ['./result-ticker.component.css']
})
export class ResultTickerComponent implements OnInit {
  draftPicks = DRAFT_PICKS;

  constructor() { }

  ngOnInit() {
  }

}
