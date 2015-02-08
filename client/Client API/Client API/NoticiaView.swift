//
//  Categoria.swift
//  Client API
//
//  Created by Andrés Pizá on 21/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class NoticiaView: UIView {
    
    @IBOutlet weak var dataInici: UILabel!
    @IBOutlet weak var dataFi: UILabel!
    @IBOutlet weak var textDescription: UITextView!
    
    func setDataInici(dataInici: String){
        self.dataInici.text = dataInici
    }
    
    func setDataFi(dataFi: String){
        self.dataFi.text = dataFi
    }
    
    func setTextDescription(text: String) {
        self.textDescription.text = text
    }
}
