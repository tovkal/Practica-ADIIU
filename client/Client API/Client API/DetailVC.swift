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

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    private func loadDetailView() {
        switch (operation!) {
        case "Categorias":
            var view = Categoria.loadFromNibNamed("CategoriaView") as Categoria
            view.setImage(data["imagen"].string!)
            view.setTitle(data["nombre"].string!)
            view.setTextDescription(data["texto"].string!)
            self.view = view
            break;
        default:
            break;
        }
    }
}
