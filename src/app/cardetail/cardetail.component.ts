import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { NavComponent } from '../nav/nav.component';


@Component({
  selector: 'app-cardetail',
  standalone: true,
  imports: [CommonModule, RouterModule, NavComponent],
  templateUrl: './cardetail.component.html',
  styleUrl: './cardetail.component.css'
})
export class CardetailComponent {

}
