//
//  Categoria.swift
//  Client API
//
//  Created by Andrés Pizá on 21/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class FarmaciaView: UIView {
    
    @IBOutlet weak var nick: UILabel!
    @IBOutlet weak var password: UILabel!
    @IBOutlet weak var level: UILabel!
    
    func setNick(nick: String){
        self.nick.text = nick
    }
    
    func setPassword(password: String){
        self.password.text = password
    }
    
    func setLevel(level: String) {
        self.level.text = level
    }
}
