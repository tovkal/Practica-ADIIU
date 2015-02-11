//
//  Categoria.swift
//  Client API
//
//  Created by Andrés Pizá on 21/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class CategoriaView: UIView {
    @IBOutlet weak var imageView: UIImageView!
    @IBOutlet weak var title: UILabel!
    @IBOutlet weak var textDescription: UITextView!
        
    func setImage(imageName: String) {
        if imageName != "" {
            // If the image does not exist, we need to download it
            var imgURL: NSURL = NSURL(string: GlobalConstants.assetEndpoint + imageName)!
            
            // Download an NSData representation of the image at the URL
            let request: NSURLRequest = NSURLRequest(URL: imgURL)

            NSURLConnection.sendAsynchronousRequest(request, queue: NSOperationQueue.mainQueue(), completionHandler: {(response: NSURLResponse!,data: NSData!,error: NSError!) -> Void in
                if !(error? != nil) {
                    var image = UIImage(data: data)
                    
                    // Store the image in to our cache
                    dispatch_async(dispatch_get_main_queue(), {
                        self.imageView.image = image
                    })
                }
                else {
                    println("Error: \(error.localizedDescription)")
                }
            })
        }
    }
    
    func setTitle(title: String){
        self.title.text = title
    }
    
    func setTextDescription(text: String) {
        self.textDescription.text = text
    }
}
