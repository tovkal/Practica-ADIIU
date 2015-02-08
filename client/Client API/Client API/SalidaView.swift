//
//  SalidaView.swift
//  Client API
//
//  Created by Andrés Pizá on 8/2/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class SalidaView: UIView {

    @IBOutlet weak var title: UILabel!
    @IBOutlet weak var quantity: UILabel!
    @IBOutlet weak var date: UILabel!
    
    func setTitle(title: String){
        self.title.text = title
    }
    
    func setQuantity(quantity: String) {
        self.quantity.text = quantity
    }
    
    func setDate(date: String) {
        self.date.text = date
    }

}
