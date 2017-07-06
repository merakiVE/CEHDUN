import { Component } from '@angular/core';

@Component({
    selector: 'api',
    templateUrl: './api.component.html'
})
export class ApiComponent{

	public title:string;

	constructor(){
		this.title = 'API';
	}

	ngOnInit(){
		console.log("Cargado el componente");
	}
}