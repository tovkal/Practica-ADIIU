//
//  OperationTVC.swift
//  Client API
//
//  Created by Andrés Pizá on 20/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class OperationTVC: UITableViewController {
    
    // API operations
    var operations = ["Categorias": "/categorias", "Entradas": "/entradas", "Salidas": "/salidas", "Medicamentos": "/medicamentos", "Noticias": "/noticias", "Farmacias": "/farmacias"]
    
    override func viewDidLoad() {
        super.viewDidLoad()
    }
    
    override func viewWillAppear(animated: Bool) {
        animateTable()
    }
    
    // MARK: - Table view data source
    
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return operations.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("cell", forIndexPath: indexPath) as UITableViewCell

        cell.textLabel?.text = Array(operations.keys)[indexPath.row]
        
        return cell
    }
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if indexPath.section == 0 {
            performSegueWithIdentifier("showList", sender: indexPath)
        }
    }
    
    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
        if let tvc = segue.destinationViewController as? ListTVC {
            if let index = sender as? NSIndexPath {
                tvc.operationTitle = Array(operations.keys)[index.row]
                tvc.method = operations[tvc.operationTitle!]
            }
        }
    }
    
    // MARK: - Table animation
    func animateTable() {
        tableView.reloadData()
        
        let cells = tableView.visibleCells()
        let tableHeight = tableView.bounds.size.height
        
        // Move all cells to the bottom of the screen
        for c in cells {
            let cell = c as UITableViewCell
            cell.transform = CGAffineTransformMakeTranslation(0, tableHeight)
        }
        
        var index = 0
        for c in cells {
            let cell = c as UITableViewCell
            UIView.animateWithDuration(1.5, delay: 0.05 * Double(index), usingSpringWithDamping: 0.8, initialSpringVelocity: 0, options: nil, animations: {
                cell.transform = CGAffineTransformMakeTranslation(0, 0)
                }, completion: nil)
            
            index++
        }
        
    }
}
