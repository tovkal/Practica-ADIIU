//
//  Medicament.swift
//  Client API
//
//  Created by Andrés Pizá on 8/2/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

extension UIView {
    class func loadFromNibNamed(nibNamed: String, bundle : NSBundle? = nil) -> UIView? {
        return UINib(
            nibName: nibNamed,
            bundle: bundle
            ).instantiateWithOwner(nil, options: nil)[0] as? UIView
    }
}