import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { Player } from '../player';
import { PLAYERS } from '../mock-players';

@Component({
  selector: 'app-players',
  templateUrl: './players.component.html',
  styleUrls: ['./players.component.css']
})
export class PlayersComponent implements OnInit {
  allPlayers = PLAYERS;
  filteredPlayers: Player[] = PLAYERS;

  filtersGroup = new FormGroup({
    showDrafted: new FormControl(false),
    showAllPositions: new FormControl('yes'),
    positionFilters: new FormGroup({
      qbFilter: new FormControl(false),
      rbFilter: new FormControl(false),
      wrFilter: new FormControl(false),
      teFilter: new FormControl(false),
      defFilter: new FormControl(false),
      kFilter: new FormControl(false)
    })
  });

  constructor() { }

  ngOnInit() {
  }

  showAllPositions() {
    console.log('showAllPositions');
    // clear individual position filter checkboxes
    let posFilters = this.filtersGroup.get('positionFilters');
    posFilters.get('qbFilter').setValue(false);
    posFilters.get('rbFilter').setValue(false);
    posFilters.get('wrFilter').setValue(false);
    posFilters.get('teFilter').setValue(false);
    posFilters.get('defFilter').setValue(false);
    posFilters.get('kFilter').setValue(false);

    // display all players
    this.filteredPlayers = this.allPlayers;

    // optionally filter out unavailable players
    if (!this.filtersGroup.get('showDrafted').value) {
      this.filteredPlayers = this.filteredPlayers.filter((p: Player) => !p.is_drafted);
    }
  }

  filterPositions() {
    console.log('filterPositions');
    let posFilters = this.filtersGroup.get('positionFilters');
    let enabledFilters = [];
    if (posFilters.get('qbFilter').value) {
        enabledFilters.push('QB');
    }
    if (posFilters.get('rbFilter').value) {
        enabledFilters.push('RB');
    }
    if (posFilters.get('wrFilter').value) {
        enabledFilters.push('WR');
    }
    if (posFilters.get('teFilter').value) {
        enabledFilters.push('TE');
    }
    if (posFilters.get('defFilter').value) {
        enabledFilters.push('DEF');
    }
    if (posFilters.get('kFilter').value) {
        enabledFilters.push('K');
    }

    if (enabledFilters.length > 0) {
      // uncheck All positions radio button
      this.filtersGroup.get('showAllPositions').setValue('no');

      // display filtered players
      this.filteredPlayers = this.allPlayers.filter((p: Player) => enabledFilters.includes(p.position));
    } else {
      // check All positions radio button
      this.filtersGroup.get('showAllPositions').setValue('yes');

      // display all players
      this.filteredPlayers = this.allPlayers;
    }

    // optionally filter out unavailable players
    if (!this.filtersGroup.get('showDrafted').value) {
      this.filteredPlayers = this.filteredPlayers.filter((p: Player) => !p.is_drafted);
    }
  }

  //getTierColor(tier: number) {
  getBackgroundColor(player: Player) {
    return '#98FB98'
  }

}
