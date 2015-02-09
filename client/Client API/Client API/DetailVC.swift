//
//  DetailVC.swift
//  Client API
//
//  Created by Andrés Pizá on 21/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit
import SwiftyJSON

class DetailVC: UIViewController {
    
    var data: JSON = JSON.nullJSON
    var operation: String?

    override func viewDidLoad() {
        super.viewDidLoad()
        
        loadDetailView()
        
        // Do not render behind navigation bar
        self.navigationController?.navigationBar.translucent = false
    }
    
    private func loadDetailView() {
        switch (operation!) {
        case "Categorias":
            var view = CategoriaView.loadFromNibNamed("CategoriaView") as CategoriaView
            view.setImage(data["imagen"].string!)
            view.setTitle(data["nombre"].string!)
            view.setTextDescription(data["texto"].string!)
            self.view = view
            break;
        case "Entradas":
            var view = EntradaView.loadFromNibNamed("EntradaView") as EntradaView
            view.setTitle(data["nombremedicamento"].string!)
            view.setQuantity(data["cantidad"].string!)
            view.setDate(data["fechahora"].string!)
            self.view = view
            break;
        case "Salidas":
            var view = SalidaView.loadFromNibNamed("SalidaView") as SalidaView
            view.setTitle(data["nombremedicamento"].string!)
            view.setQuantity(data["cantidad"].string!)
            view.setDate(data["fechahora"].string!)
            self.view = view
            break;
        case "Medicamentos":
            var view = MedicamentView.loadFromNibNamed("MedicamentView") as MedicamentView
            view.setImage(data["imagen"].string!)
            view.setTitle(data["nombre"].string!)
            view.setStock((data["enalmacen"].number!).stringValue)
            view.setTextDescription(data["texto"].string!)
            self.view = view
            break;
        case "Noticias":
            var view = NoticiaView.loadFromNibNamed("NoticiaView") as NoticiaView
            view.setDataInici(data["inicio"].string!)
            view.setDataFi(data["fin"].string!)
            view.setTextDescription(data["texto"].string!)
            self.view = view
            break;
        case "Farmacias":
            var view = FarmaciaView.loadFromNibNamed("FarmaciaView") as FarmaciaView
            view.setNick(data["nik"].string!)
            view.setPassword(data["pass"].string!)
            view.setLevel(data["nivel"].string!)
            self.view = view
            break;
        default:
            break;
        }
    }
}
