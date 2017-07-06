import { Component } from '@angular/core';

@Component({
    selector: 'home',
    templateUrl: './home.component.html'
})
export class HomeComponent{

	public title:string;

	constructor(){
		this.title = 'CEHDUN';
	}

	ngOnInit(){
		console.log("Cargado el componenete");
	}
}